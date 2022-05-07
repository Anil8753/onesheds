import { Injectable } from '@angular/core';
import { WarehouseDetail, IDetailService, QnA } from '../interfaces/detail';

@Injectable({
  providedIn: 'root',
})
export class RDetailService implements IDetailService {
  //
  constructor() {}
  getQnA(id: string, question: string): Promise<QnA[]> {
    throw new Error('Method not implemented.');
  }
  get(id: string): Promise<WarehouseDetail> {
    throw new Error('Method not implemented.');
  }
}
