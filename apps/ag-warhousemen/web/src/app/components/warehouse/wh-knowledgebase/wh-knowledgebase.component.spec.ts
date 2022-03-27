import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WhKnowledgebaseComponent } from './wh-knowledgebase.component';

describe('WhKnowledgebaseComponent', () => {
  let component: WhKnowledgebaseComponent;
  let fixture: ComponentFixture<WhKnowledgebaseComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ WhKnowledgebaseComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WhKnowledgebaseComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
