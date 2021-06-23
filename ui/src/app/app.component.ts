import { Component } from '@angular/core';
import { NzIconService } from 'ng-zorro-antd/icon';
import { environment } from 'src/environments/environment';
import { FilesService } from './services/files.service';

@Component({
  selector: 'm-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  isCollapsed = false;
  constructor(
    private filesService: FilesService,
    private NzIconService: NzIconService
  ) {
    if (environment.production) {
      NzIconService.changeAssetsSource('/app');
    }
  }
  files$ = this.filesService.list('/');
}
