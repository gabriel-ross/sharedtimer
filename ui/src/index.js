import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import Timers from './App';
import Settings from "./pages/settings";
import Timer from "./pages/timer";
import TimerSettings from "./pages/timerSettings";
import reportWebVitals from './reportWebVitals';
import { BrowserRouter, Routes, Route } from "react-router-dom";

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
        <Routes>
            <Route path="/" element={<Timers />} />
            <Route path="/timers/settings" element={<Settings />} />
            <Route path="/timers/:id" element={<Timer />} />
            <Route path="/timers/:id/settings" element={<TimerSettings />} />
        </Routes>
    </BrowserRouter>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
