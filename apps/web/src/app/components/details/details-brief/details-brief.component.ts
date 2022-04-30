import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-details-brief',
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

  constructor() {}

  ngOnInit(): void {}
}
