import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailsTopComponent } from './details-top.component';

describe('DetailsTopComponent', () => {
  let component: DetailsTopComponent;
  let fixture: ComponentFixture<DetailsTopComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DetailsTopComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DetailsTopComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
