import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import {
  WarehouseDataService,
  WarehouseItem,
} from 'src/app/services/warehouse-data.service';

@Component({
  selector: 'app-listing',
  templateUrl: './listing.component.html',
  styleUrls: ['./listing.component.scss'],
})
export class ListingComponent implements OnInit {
  //
  locality = '';
  data: WarehouseItem[] | undefined;

  constructor(
    private _route: ActivatedRoute,
    private _whDataService: WarehouseDataService
  ) {}

  ngOnInit(): void {
    this._route.queryParams.subscribe((_params) => {
      this.locality = _params['locality'];

      setTimeout(() => {
        this._whDataService.getWareHouses(this.locality).then((items) => {
          this.data = items;
        });
      }, 1000);
    });
  }
}
