import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailsMidComponent } from './details-mid.component';

describe('DetailsMidComponent', () => {
  let component: DetailsMidComponent;
  let fixture: ComponentFixture<DetailsMidComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DetailsMidComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DetailsMidComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
