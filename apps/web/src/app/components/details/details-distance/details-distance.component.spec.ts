import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailsDistanceComponent } from './details-distance.component';

describe('DetailsDistanceComponent', () => {
  let component: DetailsDistanceComponent;
  let fixture: ComponentFixture<DetailsDistanceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DetailsDistanceComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DetailsDistanceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
