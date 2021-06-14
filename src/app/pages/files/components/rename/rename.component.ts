import { Component, ElementRef, Input, OnInit, ViewChild } from '@angular/core';
import { NzModalRef } from 'ng-zorro-antd/modal';
import { FileInfo } from 'src/app/services/files.service';

@Component({
  selector: 'm-rename',
  templateUrl: './rename.component.html',
  styleUrls: ['./rename.component.scss'],
})
export class RenameComponent implements OnInit {
  constructor(private modal: NzModalRef) {}
  @ViewChild('inputRef', { static: true })
  inputRef?: ElementRef<HTMLInputElement>;
  @Input('name')
  name?: string;
  ngOnInit(): void {}
  focus() {
    console.log('focus', this.inputRef);
    this.inputRef?.nativeElement.focus();
    setTimeout(() => {});
  }
  submit() {
    this.modal.triggerOk();
  }
}
