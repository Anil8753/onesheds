import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WhTermsConditionsComponent } from './wh-terms-conditions.component';

describe('WhTermsConditionsComponent', () => {
  let component: WhTermsConditionsComponent;
  let fixture: ComponentFixture<WhTermsConditionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ WhTermsConditionsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WhTermsConditionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
