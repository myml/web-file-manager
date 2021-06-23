import { ActivatedRoute } from '@angular/router';
import { Component, HostListener, OnInit, ViewChild } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { catchError, first, map, shareReplay, switchMap } from 'rxjs/operators';
import { FileInfo, FilesService } from 'src/app/services/files.service';
import { MenuComponent } from '../menu/menu.component';
import { NzModalService } from 'ng-zorro-antd/modal';
import { MenuService } from '../../services/menu.service';

@Component({
  selector: 'm-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss'],
})
export class ListComponent implements OnInit {
  constructor(
    private route: ActivatedRoute,
    private filesService: FilesService,
    private menuService: MenuService,
    private modalService: NzModalService
  ) {}
  @ViewChild('ListNemuRef', { static: true })
  ListMenuEle!: MenuComponent;

  private selectInfo?: FileInfo;
  private refresh$ = new BehaviorSubject(null);

  workdir$ = this.route.url.pipe(
    switchMap((url) => this.refresh$.pipe(map(() => url))),
    map((url) => {
      return url.map(String).map(decodeURIComponent).join('/');
    })
  );
  files$ = this.workdir$.pipe(
    switchMap((path) => {
      return this.filesService.list(path);
    })
  );
  ngOnInit() {
    console.log('list component');
  }
  refresh() {
    this.refresh$.next(null);
  }
  async drop(ev: Event, info: FileInfo) {
    ev.preventDefault();
    if (!this.selectInfo || this.selectInfo.fullname === info.fullname) {
      return;
    }
    await this.filesService
      .move(
        this.selectInfo.fullname,
        info.fullname + '/' + this.selectInfo.name
      )
      .toPromise();
    this.refresh();
  }
  allowDrop(ev: Event) {
    ev.preventDefault();
  }
  dragstart(ev: Event, info: FileInfo) {
    this.selectInfo = info;
  }

  @HostListener('contextmenu', ['$event'])
  listMenu(ev: MouseEvent) {
    ev.preventDefault();
    ev.stopPropagation();
    if (!this.ListMenuEle) {
      return;
    }
    this.ListMenuEle.items = [
      { text: '刷新', click: () => this.refresh() },
      {
        text: '上传',
        click: async () => {
          const workdir = await this.workdir$.pipe(first()).toPromise();
          await this.menuService.upload(workdir);
          this.refresh();
        },
      },
      { text: '全选', disable: true },
      {
        text: '粘贴',
        disable: this.menuService.getClipboard() === '',
        click: async () => {
          const workdir = await this.workdir$.pipe(first()).toPromise();
          const name = this.filesService.baseName(
            this.menuService.getClipboard()
          );
          await this.filesService
            .move(this.menuService.getClipboard(), workdir + '/' + name)
            .toPromise();
          this.refresh();
        },
      },
      {
        text: '新建文件夹',
        click: async () => {
          const workdir = await this.workdir$.pipe(first()).toPromise();
          await this.menuService.mkdir(workdir);
          this.refresh();
        },
      },
    ];
    this.ListMenuEle.showToEvent(ev);
  }
}
