import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SearchCtrlComponent } from './search-ctrl.component';

describe('SearchCtrlComponent', () => {
  let component: SearchCtrlComponent;
  let fixture: ComponentFixture<SearchCtrlComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SearchCtrlComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SearchCtrlComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
