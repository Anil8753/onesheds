import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { DataService } from 'src/app/services/data.service';
import { ListItem } from 'src/app/services/interfaces/list';

@Component({
  selector: 'app-listing',
  templateUrl: './listing.component.html',
  styleUrls: ['./listing.component.scss'],
})
export class ListingComponent implements OnInit {
  //
  locality = '';
  data: ListItem[] | undefined;

  constructor(
    private _route: ActivatedRoute,
    private _dataservice: DataService
  ) {}

  ngOnInit(): void {
    this._route.queryParams.subscribe((_params) => {
      this.locality = _params['locality'];

      setTimeout(() => {
        this._dataservice.listItemService
          .get(this.locality, 10)
          .then((items) => {
            this.data = items;
          });
      }, 1000);
    });
  }
}
