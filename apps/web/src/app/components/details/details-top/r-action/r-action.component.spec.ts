import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RActionComponent } from './r-action.component';

describe('RActionComponent', () => {
  let component: RActionComponent;
  let fixture: ComponentFixture<RActionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RActionComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RActionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
