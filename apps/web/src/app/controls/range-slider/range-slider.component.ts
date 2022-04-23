import { Component, Input, OnInit } from '@angular/core';
import { Options } from '@angular-slider/ngx-slider';

@Component({
  selector: 'app-range-slider',
  templateUrl: './range-slider.component.html',
  styleUrls: ['./range-slider.component.scss'],
})
export class RangeSliderComponent implements OnInit {
  @Input() unit: string | undefined;

  minValue: number = 100;
  maxValue: number = 400;
  options: Options = {
    floor: 0,
    ceil: 500,
    translate: (value: number): string => {
      return this.unit ? this.unit + value : value.toString();
    },
    combineLabels: (minValue: string, maxValue: string): string => {
      return minValue + ' - ' + maxValue;
    },
  };

  constructor() {}

  ngOnInit(): void {}
}
