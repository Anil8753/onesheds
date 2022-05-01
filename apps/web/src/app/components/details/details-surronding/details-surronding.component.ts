import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-details-surronding',
  templateUrl: './details-surronding.component.html',
  styleUrls: ['./details-surronding.component.scss'],
})
export class DetailsSurrondingComponent implements OnInit {
  @Input() data:
    | {
        shoppingMall: string[];
        commercialArea: string[];
        ResidentialArea: string[];
        landmarks: string[];
      }
    | undefined;

  constructor() {}

  ngOnInit(): void {}
}
