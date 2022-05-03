import {
  Component,
  Input,
  OnChanges,
  OnInit,
  SimpleChanges,
} from '@angular/core';
import { WarehouseDetail } from 'src/app/services/interfaces/detail';

@Component({
  selector: 'app-details-top',
  templateUrl: './details-top.component.html',
  styleUrls: ['./details-top.component.scss'],
})
export class DetailsTopComponent implements OnInit, OnChanges {
  //
  @Input() data: WarehouseDetail | undefined;
  attractions: string[] = [];

  //
  constructor() {}
  ngOnChanges(changes: SimpleChanges): void {
    let change = changes['data'];
    if (change.previousValue != change.currentValue) {
      if (!!this.data) {
        if (this.data.available) this.attractions.push('Available');
        if (this.data.greenWarehouse) this.attractions.push('Green warehuse');
        if (this.data.insured) this.attractions.push('Insured');
        if (this.data.sharable) this.attractions.push('Sharable');
      }
    }
  }

  ngOnInit(): void {}
}
