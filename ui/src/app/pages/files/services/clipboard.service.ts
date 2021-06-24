import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ClipboardService {
  private clipboard!: Clipboard;
  constructor() {}

  setClipboard(clipboard: Clipboard) {
    this.clipboard = clipboard;
  }
  getClipboard() {
    return this.clipboard;
  }
}

interface Clipboard {
  type: 'file' | 'text';
  active: 'copy' | 'cut';
  value: string;
}
