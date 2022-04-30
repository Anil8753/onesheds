import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailsBriefComponent } from './details-brief.component';

describe('DetailsBriefComponent', () => {
  let component: DetailsBriefComponent;
  let fixture: ComponentFixture<DetailsBriefComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DetailsBriefComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DetailsBriefComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
