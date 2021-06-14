import {
  Component,
  EventEmitter,
  HostListener,
  Input,
  OnInit,
  Output,
  ViewChild,
} from '@angular/core';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';
import { NzModalService } from 'ng-zorro-antd/modal';
import { FileInfo, FilesService } from 'src/app/services/files.service';
import { MenuService } from '../../services/menu.service';
import { MenuComponent } from '../menu/menu.component';

@Component({
  selector: 'm-file',
  templateUrl: './file.component.html',
  styleUrls: ['./file.component.scss'],
})
export class FileComponent implements OnInit {
  constructor(
    private filesService: FilesService,
    private menuService: MenuService,
    private safe: DomSanitizer
  ) {}
  @ViewChild('MemuRef', { static: true })
  MenuEle!: MenuComponent;
  @Input('info')
  info!: FileInfo;
  @Output()
  change = new EventEmitter();

  previewURL!: SafeUrl;

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
        text: '下载',
        click: () => {
          const a = document.createElement('a');
          a.href = `/api/dl/${this.info.name}?path=` + this.info.fullname;
          a.download = this.info.name;
          a.click();
        },
      },
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

  preview() {
    this.previewURL = this.safe.bypassSecurityTrustResourceUrl(
      '/api/file?path=' + this.info.fullname
    );
  }

  fileIcon(name: string, ext: string) {
    if (['.json', '.ts', '.html', '.go', '.js'].includes(ext)) {
      return 'file-text';
    }
    if (['.png', '.svg', '.gif', '.jpg', '.jpeg'].includes(ext)) {
      return 'file-image';
    }
    if (['.md'].includes(ext)) {
      return 'file-markdown';
    }
    return 'file-unknown';
  }
}
