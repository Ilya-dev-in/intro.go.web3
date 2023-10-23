import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export @Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.less']
})
class AppComponent {
  title = 'webapp';

  constructor(private http: HttpClient) {}

  ngOnInit() {        
    this.http.get<ToolsDTO>("/api/root").subscribe(data => {
      this.title = data.Name
    })
  }
}

class ToolsDTO {
  Name: string
}