import { Component, OnInit } from '@angular/core';
import { HttpClient} from "@angular/common/http";


export class Error {
  constructor(
    public Success: boolean,
    public Message: string,
  ) {
  }
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit {
  err: Error | undefined
  url = 'http://localhost:8008/'
  constructor(
    private httpClient: HttpClient
  ) { }

  ngOnInit(): void {
    this.getResponse()
  }

  getResponse(){
    this.httpClient.get<any>(this.url).subscribe(
      response => {
        console.log(response);
      }
    )
  }

}
