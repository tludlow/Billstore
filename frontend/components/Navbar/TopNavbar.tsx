import { Disclosure } from '@headlessui/react'
import { ChevronDownIcon } from '@heroicons/react/outline'

export default function TopNavbar() {
    return (
        <Disclosure as="nav" className="py-1.5 text-sm text-white bg-cool-gray-800">
            {({ open }) => (
                <>
                    <div className="container flex items-center justify-between px-4 mx-auto sm:px-0">
                        <div className="flex items-center space-x-1">
                            <span>GBP</span> <ChevronDownIcon className="w-3 h-3" />
                        </div>
                        <div className="">Get free delivery on orders over £30</div>
                        <div className="flex items-center">
                            <a href="">Create an account</a>
                            <div className="w-px h-full mx-5 border-l border-gray-400">
                                <span className="text-transparent">1</span>
                            </div>
                            <a href="">Sign in</a>
                        </div>
                    </div>
                </>
            )}
        </Disclosure>
    )
}
