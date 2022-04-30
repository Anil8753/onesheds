import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailsMidActionsComponent } from './details-mid-actions.component';

describe('DetailsMidActionsComponent', () => {
  let component: DetailsMidActionsComponent;
  let fixture: ComponentFixture<DetailsMidActionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DetailsMidActionsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DetailsMidActionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
