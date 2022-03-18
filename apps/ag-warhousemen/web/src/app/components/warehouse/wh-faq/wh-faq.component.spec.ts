import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WhFaqComponent } from './wh-faq.component';

describe('WhFaqComponent', () => {
  let component: WhFaqComponent;
  let fixture: ComponentFixture<WhFaqComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ WhFaqComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WhFaqComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
