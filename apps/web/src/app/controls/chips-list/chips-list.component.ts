import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-chips-list',
  templateUrl: './chips-list.component.html',
  styleUrls: ['./chips-list.component.scss'],
})
export class ChipsListComponent implements OnInit {
  //
  @Input() heading: string;
  @Input() items: string[];
  constructor() {
    this.heading = '';
    this.items = [];
  }

  ngOnInit(): void {}
}
