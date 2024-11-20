import { SiGithub, SiLinkedin } from "@icons-pack/react-simple-icons";

function Footer() {
    return (
        <footer className="w-full border-t border-gray-200 bg-white py-2">
            <div className="container mx-auto flex flex-col-reverse items-center justify-between px-4 sm:flex-row">
                <div className="mt-2 flex space-x-4 sm:mt-0">
                    <a href="https://github.com/jamesjohnson88" className="text-gray-400 hover:text-gray-500">
                        <span className="sr-only">GitHub</span>
                        <SiGithub className="h-5 w-5"/>
                    </a>
                    <a href="https://www.linkedin.com/in/jamesjohnson88/" className="text-gray-400 hover:text-gray-500">
                        <span className="sr-only">LinkedIn</span>
                        <SiLinkedin className="h-5 w-5"/>
                    </a>
                </div>
                <p className="text-sm text-gray-500 mb-2 sm:mb-0" style={{
                    fontFamily: '"Coming Soon", cursive'
                }}>© 2024 go-todo-app. No rights reserved.</p>
            </div>
        </footer>
    )
}

export default Footer