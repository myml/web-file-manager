import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'm-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.scss'],
})
export class MenuComponent implements OnInit {
  constructor() {}
  @Input('items') items: Item[] = [];
  @Input('x') x = 0;
  @Input('y') y = 0;
  hide = true;
  private bgEl!: HTMLDivElement;

  ngOnInit() {}

  disableContextmenu(ev: MouseEvent) {
    ev.preventDefault();
    ev.stopPropagation();
  }
  private bg() {
    const bg = document.createElement('div');
    bg.style.position = 'absolute';
    bg.style.width = '100vw';
    bg.style.height = '100vh';
    // bg.style.backgroundColor = 'red';
    bg.style.top = '0';
    bg.style.left = '0';
    bg.style.zIndex = '9999';
    bg.addEventListener('click', () => {
      this.hidden();
    });
    bg.addEventListener('contextmenu', (e) => {
      e.preventDefault();
      this.hidden();
    });
    this.bgEl = bg;
    document.body.append(this.bgEl);
  }
  show() {
    this.hide = false;
    this.bg();
  }
  showToEvent(ev: MouseEvent) {
    this.x = (ev as any)['layerX'];
    this.y = (ev as any)['layerY'];
    this.hide = false;
    this.bg();
  }
  hidden() {
    if (this.bgEl) {
      document.body.removeChild(this.bgEl);
    }
    this.hide = true;
  }
}

export interface Item {
  text: string;
  disable?: boolean;
  click?: () => void;
}
