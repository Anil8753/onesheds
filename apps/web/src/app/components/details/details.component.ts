import { Component, Inject, OnInit } from '@angular/core';
import { DataService } from 'src/app/services/data.service';
import { WarehouseDetail } from 'src/app/services/interfaces/detail';

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss'],
})
export class DetailsComponent implements OnInit {
  //
  whId: any;
  data: WarehouseDetail | undefined;
  panelOpenState = false;
  constructor(private _dataservice: DataService) {}

  ngOnInit(): void {
    const promise = this._dataservice.detailService.get(this.whId);
    promise.then((r) => {
      this.data = r;
    });
  }
}
