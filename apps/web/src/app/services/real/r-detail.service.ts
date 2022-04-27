import { Injectable } from '@angular/core';
import { Detail, IDetailService } from '../interfaces/detail';

@Injectable({
  providedIn: 'root',
})
export class RDetailService implements IDetailService {
  //
  constructor() {}
  get(id: string): Promise<Detail> {
    throw new Error('Method not implemented.');
  }
}
