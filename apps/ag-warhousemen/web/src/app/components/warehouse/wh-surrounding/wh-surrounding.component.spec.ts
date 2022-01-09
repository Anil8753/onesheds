import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WhSurroundingComponent } from './wh-surrounding.component';

describe('WhSurroundingComponent', () => {
  let component: WhSurroundingComponent;
  let fixture: ComponentFixture<WhSurroundingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ WhSurroundingComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WhSurroundingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
