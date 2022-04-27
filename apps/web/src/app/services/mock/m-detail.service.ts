import { Injectable } from '@angular/core';
import { Detail, IDetailService } from '../interfaces/detail';

@Injectable({
  providedIn: 'root',
})
export class MDetailService implements IDetailService {
  //
  private item: any;

  constructor() {
    this.initTestData();
  }

  async get(id: string): Promise<Detail> {
    return this.item;
  }

  private initTestData() {
    this.item = { id: '12345', name: 'test name' };
  }
}
