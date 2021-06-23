import { Component, Input, OnInit } from '@angular/core';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';
import { FileInfo } from 'src/app/services/files.service';

@Component({
  selector: 'm-preview',
  templateUrl: './preview.component.html',
  styleUrls: ['./preview.component.scss'],
})
export class PreviewComponent implements OnInit {
  constructor(private safe: DomSanitizer) {}
  @Input()
  info!: FileInfo;
  previewURL!: SafeResourceUrl;
  ngOnInit() {}
}
