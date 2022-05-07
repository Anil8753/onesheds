import { ComponentFixture, TestBed } from '@angular/core/testing';

import { QNAComponent } from './q-n-a.component';

describe('QNAComponent', () => {
  let component: QNAComponent;
  let fixture: ComponentFixture<QNAComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ QNAComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(QNAComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
