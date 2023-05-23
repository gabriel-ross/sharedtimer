* back-end message should include a timestamp so that the client timer can properly initialize

# Protobuf stuff
- docs/getting started
https://protobuf.dev/getting-started/gotutorial/

- example
https://github.com/protocolbuffers/protobuf/blob/main/examples/go/cmd/add_person/add_person.go

https://github.com/golang/protobuf/blob/master/protoc-gen-go/grpc/grpc.go

https://github.com/golang/protobuf

- demo
https://www.youtube.com/watch?v=BdzYdN_Zd9Q


# React Native stuff
https://blog.logrocket.com/how-to-build-ios-apps-using-react-native/

extern crate chrono;
extern crate timer;
use std::sync::mpsc::{channel, Receiver, Sender};
use std::sync::{Arc, Mutex};
use std::{thread, time};

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

    let t = Ticker::new(chrono::Duration::seconds(1));
    thread::sleep(time::Duration::from_secs(2));
    let rx = t.tick();
    for x in rx {
        println!("hello");
    }
    let foo = rx.recv();
}

struct Clock(i8, i8, i8);

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

    fn tick(mut self) -> Receiver<bool> {
        let tx = self.tx.clone();
        let guard = self.timer.schedule_repeating(self.duration, move || {
            let _ignored = tx.send(true);
        });
        self.guard = Some(guard);
        self.rx
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
