import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-details-distance',
  templateUrl: './details-distance.component.html',
  styleUrls: ['./details-distance.component.scss'],
})
export class DetailsDistanceComponent implements OnInit {
  //
  @Input() data:
    | {
        important: {
          policeStation: number;
          fireStation: number;
          weighbridge: number;
          highway: number;
        };
      }
    | undefined;

  constructor() {}

  ngOnInit(): void {}
}
