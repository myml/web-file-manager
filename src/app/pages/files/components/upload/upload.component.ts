import { Component, Input, OnInit } from '@angular/core';
import { FilesService } from 'src/app/services/files.service';

@Component({
  selector: 'm-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.scss'],
})
export class UploadComponent implements OnInit {
  constructor(private filesService: FilesService) {}
  @Input()
  path?: string;
  files?: Item[];
  ngOnInit() {}

  select() {
    console.log('select upload', this.path);
    const path = this.path;
    const el = document.createElement('input');
    el.name = 'upload';
    el.type = 'file';
    el.multiple = true;
    el.addEventListener('change', async () => {
      if (!el.files || el.files.length == 0) {
        return;
      }
      this.files = Array.from(el.files).map((file) => {
        return { name: file.name, raw: file, progress: 0, status: 'wait' };
      });
      for (const item of this.files) {
        item.status = 'uploading';
        await this.filesService.upload(path || '/', item.raw, (progress) => {
          console.log(progress);
          item.progress = Math.round(progress);
        });
        item.status = 'success';
      }
    });
    el.click();
  }
}

interface Item {
  name: string;
  raw: File;
  progress: number;
  status: 'wait' | 'uploading' | 'success' | 'failure';
}
