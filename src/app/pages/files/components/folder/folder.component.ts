import {
  Component,
  EventEmitter,
  HostListener,
  Input,
  OnInit,
  Output,
  ViewChild,
} from '@angular/core';
import { FileInfo, FilesService } from 'src/app/services/files.service';
import { MenuService } from '../../services/menu.service';
import { MenuComponent } from '../menu/menu.component';

@Component({
  selector: 'm-folder',
  templateUrl: './folder.component.html',
  styleUrls: ['./folder.component.scss'],
})
export class FolderComponent implements OnInit {
  constructor(
    private filesService: FilesService,
    private menuService: MenuService
  ) {}

  @ViewChild('MemuRef', { static: true })
  private MenuEle: MenuComponent | undefined;
  @Input('info')
  info!: FileInfo;
  @Output('change')
  change = new EventEmitter();

  ngOnInit() {}
  @HostListener('contextmenu', ['$event'])
  itemMenu(ev: MouseEvent) {
    ev.preventDefault();
    ev.stopPropagation();
    if (!this.MenuEle) {
      return;
    }
    this.MenuEle.items = [
      {
        text: '剪贴',
        click: () => {
          this.menuService.setClipboard(this.info.fullname);
        },
      },
      {
        text: '复制',
        disable: true,
        click: () => {
          this.menuService.setClipboard(this.info.fullname);
        },
      },
      {
        text: '删除',
        click: async () => {
          await this.filesService.delete(this.info.fullname).toPromise();
          this.change.next();
        },
      },
      {
        text: '重命名',
        click: async () => {
          await this.menuService.rename(this.info);
          this.change.next();
        },
      },
      { text: '属性', disable: true },
    ];
    this.MenuEle.showToEvent(ev);
  }

  disableClick(ev: Event) {
    ev.preventDefault();
    ev.stopPropagation();
  }
}
