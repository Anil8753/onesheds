import { TestBed } from '@angular/core/testing';

import { WarehouseDataService } from './warehouse-data.service';

describe('WarehouseDataService', () => {
  let service: WarehouseDataService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(WarehouseDataService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
