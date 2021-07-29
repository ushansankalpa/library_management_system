import { Component, OnInit } from '@angular/core';
import {HttpClient} from '@angular/common/http'


interface IBooks{
  book_title : string
  author : string
  price : number
  available : string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{
  title = 'Library System';

  public book_id = ''
  public book_title = ''
  public author = ''
  public price =0
  public available = ''


  public books : IBooks[] = []

  constructor(
    private httpClient : HttpClient
  ){}

  async ngOnInit(){
    await this.loadBooks()
  }

  async loadBooks(){
    this.books = await this.httpClient.get<IBooks[]>('/api/books/get').toPromise()
  }

  async addBook(){
    await this.httpClient.post('api/books/add', {
      //book_id :  Math.random(),
      book_title: this.book_title,
      author: this.author,
      price: this.price,
      available : this.available
    }).toPromise()
    // this.books.push({
    //   book_title: this.book_title,
    //   author: this.author,
    //   price: this.price,
    //   available : this.available
    // })
    
    this.book_title = ''
    this.author = ''
    this.price=0
    this.available = ''
    await this.loadBooks()
  }
}
