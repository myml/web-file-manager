import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzListModule } from 'ng-zorro-antd/list';
import { NzProgressModule } from 'ng-zorro-antd/progress';

import { FilesRoutingModule } from './files-routing.module';
import { ListComponent } from './components/list/list.component';
import { FolderComponent } from './components/folder/folder.component';
import { FileComponent } from './components/file/file.component';
import { MenuComponent } from './components/menu/menu.component';
import { RenameComponent } from './components/rename/rename.component';
import { MenuService } from './services/menu.service';
import { UploadComponent } from './components/upload/upload.component';

@NgModule({
  declarations: [
    ListComponent,
    FolderComponent,
    FileComponent,
    MenuComponent,
    RenameComponent,
    UploadComponent,
  ],
  imports: [
    CommonModule,
    FilesRoutingModule,
    NzIconModule,
    DragDropModule,
    HttpClientModule,
    NzMenuModule,
    NzModalModule,
    NzInputModule,
    NzButtonModule,
    NzListModule,
    NzProgressModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  providers: [MenuService],
})
export class FilesModule {}
