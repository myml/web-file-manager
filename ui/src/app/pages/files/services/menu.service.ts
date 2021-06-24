import { Injectable } from '@angular/core';
import { NzModalService } from 'ng-zorro-antd/modal';
import { FileInfo, FilesService } from 'src/app/services/files.service';
import { RenameComponent } from '../components/rename/rename.component';
import { UploadComponent } from '../components/upload/upload.component';

@Injectable()
export class MenuService {
  constructor(
    private modal: NzModalService,
    private filesService: FilesService
  ) {}
  private clipboard = '';
  mkdir(workdir: string) {
    return new Promise((resolve) => {
      const m = this.modal.create({
        nzTitle: '新建文件夹',
        nzAutofocus: null,
        nzContent: RenameComponent,
        nzComponentParams: { name: '新建文件夹' },
        nzOnOk: async () => {
          await this.filesService
            .mkdir(workdir + '/' + m.componentInstance?.name)
            .toPromise();
        },
      });
      m.afterOpen.subscribe(() => m.componentInstance?.focus());
      m.afterClose.subscribe(() => resolve(null));
    });
  }
  rename(info: FileInfo) {
    return new Promise((resolve) => {
      const m = this.modal.create({
        nzTitle: '重命名',
        nzAutofocus: null,
        nzContent: RenameComponent,
        nzComponentParams: { name: info.name },
        nzOnOk: async () => {
          console.log('ok');
          const fullname = info.fullname;
          const new_path =
            this.filesService.baseDir(fullname) +
            '/' +
            m.componentInstance?.name;
          await this.filesService.move(info.fullname, new_path).toPromise();
        },
      });
      m.afterOpen.subscribe(() => m.componentInstance?.focus());
      m.afterClose.subscribe(() => resolve(null));
    });
  }
  upload(workdir: string) {
    return new Promise((resolve) => {
      const m = this.modal.create({
        nzTitle: '上传',
        nzAutofocus: null,
        nzContent: UploadComponent,
        nzComponentParams: { path: workdir },
      });
      m.afterClose.subscribe(() => resolve(null));
    });
  }
}
