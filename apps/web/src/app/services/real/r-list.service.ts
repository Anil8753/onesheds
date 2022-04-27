import { Injectable } from '@angular/core';
import { IItemListService, ListItem } from '../interfaces/list';

@Injectable({
  providedIn: 'root',
})
export class RListService implements IItemListService {
  //
  constructor() {}
  get(location: string, km: number): Promise<ListItem[]> {
    throw new Error('Method not implemented.');
  }
}
