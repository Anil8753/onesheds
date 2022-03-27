import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WhReviewComponent } from './wh-review.component';

describe('WhReviewComponent', () => {
  let component: WhReviewComponent;
  let fixture: ComponentFixture<WhReviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ WhReviewComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WhReviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
