import { Injectable } from '@angular/core';
import {IUserCredentialsDefault} from '../interface/user';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {environment} from '../../../environments/environment';

@Injectable({
  providedIn: 'root'
})

export class UserService {

  constructor(
    private http: HttpClient
  ) { }

  public postUser({username, password}: IUserCredentialsDefault) {
    return this.http.post(`${environment.apiUrl}/signup`, {username, password}, {
      headers: new HttpHeaders({
        "Content-Type": "text/plain"
      })
    });
  }
}
