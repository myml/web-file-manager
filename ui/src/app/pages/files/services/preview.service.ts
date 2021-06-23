import { Injectable } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';
import { NzModalService } from 'ng-zorro-antd/modal';
import { FileInfo, FilesService } from 'src/app/services/files.service';
import { PreviewComponent } from '../components/preview/preview.component';

@Injectable()
export class PreviewService {
  constructor(
    private filesService: FilesService,
    private safe: DomSanitizer,
    private modal: NzModalService
  ) {}
  preview(info: FileInfo) {
    const url = this.safe.bypassSecurityTrustResourceUrl(
      '/api/file?path=' + info.fullname
    );
    this.modal.create({
      nzTitle: '预览',
      nzAutofocus: null,
      nzContent: PreviewComponent,
      nzComponentParams: { info: info, previewURL: url },
      nzFooter: null,
    });
  }
}
