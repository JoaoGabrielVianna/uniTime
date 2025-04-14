import { LoginForm } from "@/components/login-form"

export default function LoginPage() {
  return (
    <div className="grid min-h-svh lg:grid-cols-2">
      <div className="flex flex-col gap-4 p-6 md:p-10">
        <div className="flex justify-center gap-2 md:justify-start">
          <a href="#" className="flex items-center font-medium relative">
          <span className="font-medium">Uni</span>
              <span className="text-yellow-600">Time</span>
              <div className="absolute -top-1 -right-3 w-2 h-2 bg-yellow-600 rounded-full animate-pulse"></div>
          </a>
        </div>
        <div className="flex flex-1 items-center justify-center">
          <div className="w-full max-w-xs">
            <LoginForm />
          </div>
        </div>
      </div>
      <div className="relative hidden bg-muted lg:block">
        
      </div>
    </div>
  )
}
