import React, { useEffect, useState } from 'react';
import Navbar from './components/Navbar';
import { callAPI } from '../settings/http';

interface UserProfile {
  name: string;
  role: string;
  email: string;
  joinDate: string;
  bio: string;
  skills: string[];
  mobileNumbers: string[];
  firstName: string;
  lastName: string;
  profilePictureUrl: string;
  dateOfBirth: string;
}

const Profile = () => {
  const [userProfile, setUserProfile] = useState<UserProfile | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => { 
    const fetchUserProfile = async () => {
      try {
        const response = await callAPI('/user/profile', 'GET', null, false);
        const data = await response;
        setUserProfile(data);
      } catch (err: any) {
        setError(err.message);
      } finally {
        setIsLoading(false);
      }
    };
    fetchUserProfile();
  }, []);

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar></Navbar>
    <div className="min-h-screen bg-[#000300] text-white p-4">
      <div className="max-w-4xl mx-auto">
        {isLoading ? (
          <div className="flex justify-center items-center h-96">
            <div className="animate-spin rounded-full h-16 w-16 border-t-2 border-b-2 border-[#00df9a]"></div>
          </div>
        ) : (
          <>
            {/* Profile Header */}
            <div className="bg-black rounded-xl p-6 mb-6">
              <div className="flex flex-col md:flex-row items-center gap-6">
                {/* Profile Image */}
                <div className="relative w-32 h-32">
                  {userProfile?.profilePictureUrl ? (
                    <img 
                      src={userProfile.profilePictureUrl}
                      alt="Profile"
                      className="w-full h-full rounded-full object-cover"
                    />
                  ) : (
                    <div className="w-full h-full rounded-full bg-[#00df9a] flex items-center justify-center">
                      <span className="text-4xl text-black font-bold">
                        {userProfile?.firstName?.charAt(0) || ""}
                      </span>
                    </div>
                  )}
                  <label className="absolute bottom-0 right-0 bg-[#00df9a] p-2 rounded-full cursor-pointer hover:bg-[#00bf82]">
                    <input 
                      type="file" 
                      className="hidden"
                      accept="image/*"
                      onChange={(e) => {
                        // Handle image upload
                        const file = e.target.files?.[0];
                        if (file) {
                          // Add your image upload logic here
                        }
                      }}
                    />
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z" />
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 13a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                  </label>
                </div>
                
                {/* Profile Info */}
                <div className="text-center md:text-left w-full">
                  <div className="flex items-center gap-2 mb-2">
                    <input
                      type="text"
                      value={userProfile?.firstName || ''}
                      className="bg-transparent border-b border-[#00df9a] text-3xl font-bold text-[#00df9a] focus:outline-none"
                      onChange={(e) => {
                        // Handle firstName update
                      }}
                    />
                    <input
                      type="text"
                      value={userProfile?.lastName || ''}
                      className="bg-transparent border-b border-[#00df9a] text-3xl font-bold text-[#00df9a] focus:outline-none"
                      onChange={(e) => {
                        // Handle lastName update
                      }}
                    />
                  </div>
                  <input
                    type="date"
                    value={userProfile?.dateOfBirth || ''}
                    className="bg-transparent border-b border-gray-400 text-xl text-gray-400 focus:outline-none mb-2"
                    onChange={(e) => {
                      // Handle dateOfBirth update
                    }}
                  />
                </div>
              </div>
            </div>

            {/* Profile Details */}
            <div className="grid md:grid-cols-2 gap-6">
              {/* Contact Information */}
              <div className="bg-black rounded-xl p-6 md:col-span-2">
                <h2 className="text-2xl font-bold text-[#00df9a] mb-4">Contact Information</h2>
                <div className="space-y-4">
                  <div className="flex items-center gap-2">
                    <span className="font-semibold text-gray-300">Email:</span>
                    <input
                      type="email"
                      value={userProfile?.email || ''}
                      className="bg-transparent border-b border-gray-300 text-gray-300 focus:outline-none flex-1"
                      onChange={(e) => {
                        // Handle email update
                      }}
                    />
                  </div>
                  <div>
                    <span className="font-semibold text-gray-300">Phone Numbers:</span>
                    <div className="mt-2 space-y-2">
                      {userProfile?.mobileNumbers?.map((number, index) => (
                        <div key={index} className="flex items-center gap-2">
                          <input
                            type="tel"
                            value={number}
                            className="bg-transparent border-b border-gray-300 text-gray-300 focus:outline-none flex-1 ml-4"
                            onChange={(e) => {
                              // Handle phone number update
                            }}
                          />
                          <button 
                            className="text-red-500 hover:text-red-400"
                            onClick={() => {
                              // Handle remove phone number
                            }}
                          >
                            ×
                          </button>
                        </div>
                      ))}
                      <button 
                        className="text-[#00df9a] hover:text-[#00bf82] ml-4"
                        onClick={() => {
                          // Handle add new phone number
                        }}
                      >
                        + Add Phone Number
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </>
        )}
      </div>
    </div>
    </div>
  );
};

export default Profile;