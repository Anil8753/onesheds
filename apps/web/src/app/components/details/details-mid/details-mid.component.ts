import {
  Component,
  Input,
  OnChanges,
  OnInit,
  SimpleChanges,
} from '@angular/core';
import { WarehouseDetail } from 'src/app/services/interfaces/detail';

@Component({
  selector: 'app-details-mid',
  templateUrl: './details-mid.component.html',
  styleUrls: ['./details-mid.component.scss'],
})
export class DetailsMidComponent implements OnInit, OnChanges {
  @Input() data: WarehouseDetail | undefined;
  amenities: string[] = [];
  storageOptions: string[] = [];
  accreditation: string[] = [];
  //
  constructor() {}

  ngOnChanges(changes: SimpleChanges): void {
    let change = changes['data'];
    if (change.previousValue != change.currentValue) {
      if (!this.data) return;

      this.amenities = this.data.amenities
        .filter((e) => e.value)
        .map((e, i) => {
          return e.item;
        });

      this.storageOptions = this.data.storageOptions
        .filter((e) => e.value)
        .map((e, i) => {
          return e.item;
        });

      this.accreditation = this.data.accrediatation
        .filter((e) => e.value)
        .map((e, i) => {
          return e.item;
        });
    }
  }

  ngOnInit(): void {}
}
