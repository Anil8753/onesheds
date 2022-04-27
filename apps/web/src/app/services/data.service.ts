import { Injectable } from '@angular/core';
import { IDetailService } from './interfaces/detail';
import { IItemListService } from './interfaces/list';
import { MDetailService } from './mock/m-detail.service';
import { MListService } from './mock/m-list.service';
import { RDetailService } from './real/r-detail.service';
import { RListService } from './real/r-list.service';

@Injectable({
  providedIn: 'root',
})
export class DataService {
  //
  public listItemService: IItemListService;
  public detailService: IDetailService;

  constructor() {
    if (true) {
      this.listItemService = new MListService();
      this.detailService = new MDetailService();
    } else {
      this.listItemService = new RListService();
      this.detailService = new RDetailService();
    }
  }
}
