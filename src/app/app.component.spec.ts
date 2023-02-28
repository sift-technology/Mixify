import { HttpClient } from '@angular/common/http';
import { TestBed, ComponentFixture, fakeAsync } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { Router } from "@angular/router";
import { Location } from '@angular/common';
import { AppComponent } from './app.component';
import { AppModule } from './app.module';
import { QuestionComponent } from './question/question.component';
import { ResultsComponent } from './results/results.component';


describe('AppComponent', () => {
  let router: Router;
  let location: Location;
  let app: AppComponent;
  let fixture: ComponentFixture<AppComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RouterTestingModule,
        AppModule,
        RouterTestingModule.withRoutes([
          { path: '', component: QuestionComponent },
          { path: 'results', component: ResultsComponent }
        ])
      ],
      declarations: [
        AppComponent
      ],
    }).compileComponents();

    router = TestBed.inject(Router);
    location = TestBed.inject(Location);
    fixture = TestBed.createComponent(AppComponent);
    app = fixture.componentInstance;
  });

  it('should create the app', () => {
    expect(app).toBeTruthy();
  });

  it(`title of app should be 'Mixify'`, () => {
    expect(app.title).toEqual('Mixify');
  });

  it(`title in toolbar should be 'Mixify'`, () => {
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('mat-toolbar').textContent).toContain('Mixify');
  });

  it('navigate to "results" takes you to /results', fakeAsync(() => {
    router.initialNavigation();
    router.navigate(["/results"]).then(() => {
      expect(location.path()).toBe("/results");
    });
  }));
});
