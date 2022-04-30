import { Injectable } from '@angular/core';
import { WarehouseDetail, IDetailService } from '../interfaces/detail';

@Injectable({
  providedIn: 'root',
})
export class RDetailService implements IDetailService {
  //
  constructor() {}
  get(id: string): Promise<WarehouseDetail> {
    throw new Error('Method not implemented.');
  }
}
