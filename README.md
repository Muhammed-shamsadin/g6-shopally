# ShopAlly – AI-Powered Shopping Assistant for AliExpress  
Helping customers find the right products faster with **AI search, multilingual support, and smart notifications**.  

---

## 🚀 Overview  
ShopAlly is an **AI-powered shopping platform** that assists AliExpress customers in making faster and smarter purchase decisions.  
> ⚠️ Note: The platform is **partially completed**. Many core features are implemented, but full integration and deployment are still in progress.  

It reduces the time spent searching for products by combining:  

- **AI-powered product search**  
- **Multilingual support (English & Amharic)**  
- **Redis rate-limiting** for secure API usage  
- **Price-drop notifications**  
- **CI/CD pipelines** for reliable deployments  

---

## 👨‍💻 My Contributions  
This fork highlights the parts of ShopAlly I personally worked on:  

- **Search Endpoint (/search)**  
  - Implemented the **bridge between ShopAlly and the AliExpress API**.  
  - Fetched product data from AliExpress and **remapped fields** into ShopAlly’s custom structure for AI processing.  

- **Code Quality & Reviews**  
  - Reviewed team PRs to ensure code quality and smooth merges.  

- **CI/CD**  
  - Fixed **GitHub Actions CI issues**, ensuring builds and tests ran successfully.  

---

## 🖼 UI Screenshots
Here are some UI examples of the platform (replace the paths with your actual screenshots in your repo):  

**Web UI**  
![Home Page](./screenshots/home.png)  
![Search Results](./screenshots/search.png)  

**Mobile UI**  
![Mobile Home](./screenshots/mobile-home.png)  
![Mobile Search](./screenshots/mobile-search.png)  

---

## 🛠 Tech Stack  
- **Backend:** Go (Gin), Clean Architecture  
- **Database:** MongoDB  
- **Caching & Rate Limiting:** Redis  
- **Auth:** Google OAuth  
- **CI/CD:** GitHub Actions, Docker  
- **AI Integration (Planned):** Query parsing for smarter search  
- **Languages:** English + Amharic  

---

## 📂 Proof of Contribution  
This is a **forked repository** to highlight my backend contributions.  

- 🔗 [Original Team Repository](https://github.com/A2SV/g6-shopally)  
- 👥 [Contributors Page (Proof)](https://github.com/A2SV/g6-shopally/graphs/contributors)  
- [Example Commit](https://github.com/A2SV/g6-shopally/commit/2f1a5316c0ae94e348e887a5db9a8c39fc054b7f)

---

## ⚙️ Setup & Installation  

### 1. Clone the repository  
```bash
git clone https://github.com/YOUR_USERNAME/shopally.git
cd shopally
```
### 2. Backend Setup

Navigate to the backend folder and install dependencies:
```
cd shopally-backend
go mod tidy
```

Start the backend server:
```
go run main.go
```
### 3. Frontend Setup

Navigate to the frontend folder and install dependencies:
```
cd ../shopally-frontend
pnpm install
```

Start the frontend server:
```
pnpm dev
```
### 4. Mobile App Setup (Flutter)

Navigate to the mobile app folder:
```
cd ../shopally-mobile
flutter pub get
```

Run the app on an emulator or device:
```
flutter run
```
### 5. Access the Application

Web: http://localhost:3000

Mobile: Running emulator or connected device


This is fully **copyable** and ready to paste into your `README.md`.  

If you want, I can also **add a “Project Overview” and “Technologies Used” section** so your README looks professional and complete. Do you want me to do that?

