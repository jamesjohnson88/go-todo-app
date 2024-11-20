import Header from "@/components/header/Header.tsx";
import TodoItemList from "@/components/todo-items/TodoItemList.tsx";
import Footer from "@/components/footer/Footer.tsx";

function App() {

  return (
    <>
      <div className={"main flex flex-col min-h-screen bg-gray-300"}>
          <Header />
          <div className={"main-body flex-grow"}>
              <TodoItemList />
          </div>
          <Footer />
      </div>
    </>
  )
}

export default App
