import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailsLocalityRatingComponent } from './details-locality-rating.component';

describe('DetailsLocalityRatingComponent', () => {
  let component: DetailsLocalityRatingComponent;
  let fixture: ComponentFixture<DetailsLocalityRatingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DetailsLocalityRatingComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DetailsLocalityRatingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
