import { Component, Input, OnInit } from '@angular/core';
import { trigger, style, animate, transition } from '@angular/animations';

@Component({
  selector: 'app-details-brief',
  animations: [
    trigger('enterAnimation', [
      transition(':enter', [
        style({ transform: 'translateX(1%)', opacity: 0 }),
        animate('50ms', style({ transform: 'translateX(0)', opacity: 1 })),
      ]),
      transition(':leave', [
        style({ transform: 'translateX(0)', opacity: 1 }),
        animate('50ms', style({ transform: 'translateX(1%)', opacity: 0 })),
      ]),
    ]),
  ],
  templateUrl: './details-brief.component.html',
  styleUrls: ['./details-brief.component.scss'],
})
export class DetailsBriefComponent implements OnInit {
  //
  @Input() heading: string = '';

  @Input() data: {
    item: string;
    value: boolean;
  }[] = [];

  @Input() dataStr: {
    item: string;
    value: string;
  }[] = [];

  @Input() dataArray: string[] = [];

  isOpen = true;
  arrowIcon = 'chevron_right';

  constructor() {}

  ngOnInit(): void {
    this.arrowIcon = this.isOpen ? 'expand_more' : 'chevron_right';
  }

  toggle() {
    this.isOpen = !this.isOpen;
    this.arrowIcon = this.isOpen ? 'expand_more' : 'chevron_right';
  }
}
