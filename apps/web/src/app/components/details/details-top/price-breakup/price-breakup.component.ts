import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-price-breakup',
  templateUrl: './price-breakup.component.html',
  styleUrls: ['./price-breakup.component.scss'],
})
export class PriceBreakupComponent implements OnInit {
  @Input() data: any;
  constructor() {}

  ngOnInit(): void {}
}
