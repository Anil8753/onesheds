import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailsSurrondingComponent } from './details-surronding.component';

describe('DetailsSurrondingComponent', () => {
  let component: DetailsSurrondingComponent;
  let fixture: ComponentFixture<DetailsSurrondingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DetailsSurrondingComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DetailsSurrondingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
