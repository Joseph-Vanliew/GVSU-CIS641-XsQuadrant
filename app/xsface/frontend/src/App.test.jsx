// App.test.js
import React from 'react';
import { render, screen } from '@testing-library/react';
import { BrowserRouter as Router } from 'react-router-dom';
import App from './App';

describe('App Component', () => {
  it('should render the app without crashing', () => {
    render(
      <Router>
        <App />
      </Router>
    );
    const appElement = screen.getByTestId('app-component');
    expect(appElement).toBeTruthy();
  });

  it('should have the title as "ui"', () => {
    render(
      <Router>
        <App />
      </Router>
    );
    const appElement = screen.getByTestId('app-component');
    expect(appElement.getAttribute('title')).toEqual('ui');
  });

  it('should render the title', () => {
    render(
      <Router>
        <App />
      </Router>
    );
    const titleElement = screen.getByText(/ui app is running!/i);
    expect(titleElement).toBeInTheDocument();
  });
});
