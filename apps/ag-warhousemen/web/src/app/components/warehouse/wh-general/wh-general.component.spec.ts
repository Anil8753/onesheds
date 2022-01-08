import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WhGeneralComponent } from './wh-general.component';

describe('WhGeneralComponent', () => {
  let component: WhGeneralComponent;
  let fixture: ComponentFixture<WhGeneralComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ WhGeneralComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WhGeneralComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
