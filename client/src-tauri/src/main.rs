// // Prevents additional console window on Windows in release, DO NOT REMOVE!!
// #![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

// // Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
// #[tauri::command]
// fn greet(name: &str) -> String {
//     format!("Hello, {}! You've been greeted from Rust!", name)
// }

// fn main() {
//     tauri::Builder::default()
//         .invoke_handler(tauri::generate_handler![greet])
//         .run(tauri::generate_context!())
//         .expect("error while running tauri application");
// }

extern crate chrono;
extern crate timer;
use std::sync::mpsc::{channel, Receiver, Sender};
use std::sync::{Arc, Mutex};
use std::{thread, time};
use uuid::uuid;

use timer::Guard;

fn main() {
    // let timer = timer::Timer::new();
    // let (tx, rx) = channel();
    // // Number of times the callback has been called.
    // let count = Arc::new(Mutex::new(0));

    // // Start repeating. Each callback increases `count`.
    // let guard = {
    //     timer.schedule_repeating(chrono::Duration::seconds(1), move || {
    //         let _ignored = tx.send(());
    //     })
    // };

    // // Sleep one second. The callback should be called ~200 times.
    // thread::sleep(std::time::Duration::new(1, 0));
    // let count_result = *count.lock().unwrap();
    // assert!(
    //     190 <= count_result && count_result <= 210,
    //     "The timer was called {} times",
    //     count_result
    // );

    // // Now drop the guard. This should stop the timer.
    // drop(guard);
    // thread::sleep(std::time::Duration::new(0, 100));

    // // Let's check that the count stops increasing.
    // let count_start = *count.lock().unwrap();
    // thread::sleep(std::time::Duration::new(1, 0));
    // let count_stop = *count.lock().unwrap();
    // assert_eq!(count_start, count_stop);

    let mut t = Ticker::new(chrono::Duration::seconds(1));
    // thread::sleep(time::Duration::from_secs(2));
    let rx = t.tick();

    // let stopper = thread::spawn(move || {
        
    // });
    thread::sleep(time::Duration::from_secs(2));
    println!("stopping");
    t.stop();

    // for x in rx {
    //     println!("hello");
    // }
}

struct CountdownTimer {
    id: uuid,

}

struct Ticker {
    timer: timer::Timer,
    guard: Option<timer::Guard>,
    duration: chrono::Duration,
    tx: Sender<bool>,
    rx: Receiver<bool>,
}

impl Ticker {
    pub fn new(duration: chrono::Duration) -> Self {
        let (tx, rx) = channel();
        Self {
            timer: timer::Timer::new(),
            guard: None,
            duration,
            tx: tx,
            rx: rx,
        }
    }

    fn tick(&mut self) -> &Receiver<bool> {
        let tx = self.tx.clone();
        let guard = self.timer.schedule_repeating(self.duration, move || {
            let _ignored = tx.send(true);
            println!("hello");
        });
        self.guard = Some(guard);
        &self.rx
    }

    fn stop(&mut self) {
        self.guard = None;
    }
}
// struct Ticker<'a> {
//     pub rx: Receiver<bool>,
//     guard: &'a timer::Guard,
// }

// impl<'a> Ticker<'a> {
//     pub fn new(duration: chrono::Duration) -> Self {
//         let (tx, rx) = channel();
//         Self {
//             rx: rx,
//             guard: &timer::Timer::new().schedule_repeating(duration, move || {
//                 let _ignored = tx.send(true);
//             }),
//         }
//     }
// }

// impl<'a> Drop for Ticker<'a> {
//     fn drop(&mut self) {
//         drop(&self.guard);
//     }
// }
